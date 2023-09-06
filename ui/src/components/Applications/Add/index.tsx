import {
  Button,
  Flex,
  GitRepository,
  Link,
  LoadingPage,
  useListSources,
} from '@weaveworks/weave-gitops';
import { Box } from '@material-ui/core';
import { PageRoute } from '@weaveworks/weave-gitops/ui/lib/types';
import _ from 'lodash';
import React, {
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from 'react';
import { useHistory } from 'react-router-dom';
import {
  ClusterAutomation,
  CreateAutomationsPullRequestRequest,
  RenderAutomationResponse,
  RepositoryRef,
} from '../../../cluster-services/cluster_services.pb';
import CallbackStateContextProvider from '../../../contexts/GitAuth/CallbackStateContext';
import useNotifications from '../../../contexts/Notifications';
import {
  expiredTokenNotification,
  useIsAuthenticated,
} from '../../../hooks/gitprovider';
import useProfiles from '../../../hooks/profiles';
import { ProfilesIndex } from '../../../types/custom';
import { useCallbackState } from '../../../utils/callback-state';
import { validateFormData } from '../../../utils/form';
import { getGitRepoHTTPSURL } from '../../../utils/formatters';
import { Routes } from '../../../utils/nav';
import { removeToken } from '../../../utils/request';
import { getGitRepos } from '../../Clusters';
import { clearCallbackState, getProviderToken } from '../../GitAuth/utils';
import { Page } from '../../Layout/App';
import { NotificationsWrapper } from '../../Layout/NotificationsWrapper';
import GitOps from '../../Templates/Form/Partials/GitOps';
import Preview from '../../Templates/Form/Partials/Preview';
import Profiles from '../../Templates/Form/Partials/Profiles';
import {
  FormWrapper,
  getRepositoryUrl,
  useGetInitialGitRepo,
} from '../../Templates/Form/utils';
import AppFields from './form/Partials/AppFields';
import { EnterpriseClientContext } from '../../../contexts/EnterpriseClient';

interface FormData {
  repo: GitRepository | null;
  provider: string;
  branchName: string;
  pullRequestTitle: string;
  commitMessage: string;
  pullRequestDescription: string;
  source_name: string;
  source_namespace: string;
  source: string;
  source_type: string;
  source_url: string;
  source_branch: string;
  clusterAutomations: {
    name: string;
    namespace: string;
    target_namespace: string;
    cluster_name: string;
    cluster_namespace: string;
    cluster: string;
    cluster_isControlPlane: boolean;
    createNamespace: boolean;
    path: string;
    source_name: string;
    source_namespace: string;
    source: string;
    source_type: string;
    source_url: string;
    source_branch: string;
  }[];
}

function getInitialData(
  callbackState: { state: { formData: FormData } } | null,
  random: string,
) {
  let defaultFormData = {
    repo: null,
    provider: '',
    branchName: `add-application-branch-${random}`,
    pullRequestTitle: 'Add application',
    commitMessage: 'Add application',
    pullRequestDescription: 'This PR adds a new application',
    clusterAutomations: [
      {
        name: '',
        namespace: '',
        target_namespace: '',
        cluster_name: '',
        cluster_namespace: '',
        cluster: '',
        cluster_isControlPlane: false,
        createNamespace: false,
        path: '',
        source_name: '',
        source_namespace: '',
        source: '',
        source_type: '',
        source_url: '',
        source_branch: '',
      },
    ],
  };

  const initialFormData = {
    ...defaultFormData,
    ...callbackState?.state?.formData,
  };

  return { initialFormData };
}

const AddApplication = ({ clusterName }: { clusterName?: string }) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [showAuthDialog, setShowAuthDialog] = useState(false);
  const { setNotifications } = useNotifications();
  const history = useHistory();
  const authRedirectPage = `/applications/create`;
  const [formError, setFormError] = useState<string>('');
  const { api } = useContext(EnterpriseClientContext);

  const optionUrl = (url?: string, branch?: string) => {
    const linkText = branch ? (
      <>
        {url}@<strong>{branch}</strong>
      </>
    ) : (
      url
    );
    if (branch) {
      return (
        <Link href={getGitRepoHTTPSURL(url, branch)} newTab>
          {linkText}
        </Link>
      );
    } else {
      return (
        <Link href={getGitRepoHTTPSURL(url)} newTab>
          {linkText}
        </Link>
      );
    }
  };

  const random = useMemo(() => Math.random().toString(36).substring(7), []);

  const callbackState = useCallbackState();

  const { initialFormData } = getInitialData(callbackState, random);

  const [formData, setFormData] = useState<any>(initialFormData);
  const firstAuto = formData.clusterAutomations[0];
  const helmRepo: RepositoryRef = useMemo(() => {
    return {
      name: firstAuto.source_name,
      namespace: firstAuto.source_namespace,
      cluster: {
        name: firstAuto.cluster_name,
        namespace: firstAuto.cluster_namespace,
      },
    };
  }, [firstAuto]);

  const { profiles, isLoading: profilesIsLoading } = useProfiles(
    firstAuto.source_type === 'HelmRepository',
    undefined,
    undefined,
    helmRepo,
  );
  const [updatedProfiles, setUpdatedProfiles] = useState<ProfilesIndex>({});
  const [openPreview, setOpenPreview] = useState(false);
  const [previewLoading, setPreviewLoading] = useState<boolean>(false);
  const [prPreview, setPRPreview] = useState<RenderAutomationResponse | null>(
    null,
  );
  const { data } = useListSources();
  const gitRepos = React.useMemo(
    () => getGitRepos(data?.result),
    [data?.result],
  );
  const initialGitRepo = useGetInitialGitRepo(null, gitRepos);

  useEffect(() => {
    setUpdatedProfiles({
      ..._.keyBy(profiles, 'name'),
      ...callbackState?.state?.updatedProfiles,
    });
  }, [callbackState?.state?.updatedProfiles, profiles]);

  useEffect(() => clearCallbackState(), []);

  useEffect(() => {
    setFormData((prevState: any) => ({
      ...prevState,
      pullRequestTitle: `Add application ${(formData.clusterAutomations || [])
        .map((a: any) => a.name)
        .join(', ')}`,
    }));
  }, [formData.clusterAutomations]);

  const getKustomizations = useCallback(() => {
    let clusterAutomations: ClusterAutomation[] = [];
    const selectedProfilesList = _.sortBy(
      Object.values(updatedProfiles),
      'name',
    ).filter(p => p.selected);
    if (formData.source_type === 'HelmRepository') {
      for (let kustomization of formData.clusterAutomations) {
        for (let profile of selectedProfilesList) {
          let values: string = '';
          let version: string = '';
          for (let value of profile.values) {
            if (value.selected === true) {
              version = value.version;
              values = value.yaml;
              clusterAutomations.push({
                cluster: {
                  name: kustomization.cluster_name,
                  namespace: kustomization.cluster_namespace,
                },
                isControlPlane: kustomization.cluster_isControlPlane,
                helmRelease: {
                  metadata: {
                    name: profile.name,
                    namespace: profile.namespace,
                  },
                  spec: {
                    chart: {
                      spec: {
                        chart: profile.name,
                        sourceRef: {
                          name: formData.source_name,
                          namespace: formData.source_namespace,
                        },
                        version,
                      },
                    },
                    values,
                  },
                },
              });
            }
          }
        }
      }
    } else {
      clusterAutomations = formData.clusterAutomations.map(
        (kustomization: any) => {
          return {
            cluster: {
              name: kustomization.cluster_name,
              namespace: kustomization.cluster_namespace,
            },
            isControlPlane: kustomization.cluster_isControlPlane,
            kustomization: {
              metadata: {
                name: kustomization.name,
                namespace: kustomization.namespace,
              },
              spec: {
                path: kustomization.path,
                sourceRef: {
                  name: kustomization.source_name,
                  namespace: kustomization.source_namespace,
                },
                targetNamespace: kustomization.target_namespace,
                createNamespace: kustomization.createNamespace,
              },
            },
          };
        },
      );
    }
    return clusterAutomations;
  }, [
    formData.clusterAutomations,
    formData.source_name,
    formData.source_namespace,
    formData.source_type,
    updatedProfiles,
  ]);

  useEffect(() => {
    if (!formData.repo) {
      setFormData((prevState: any) => ({
        ...prevState,
        repo: initialGitRepo,
      }));
    }
  }, [initialGitRepo, formData.repo]);

  const handlePRPreview = useCallback(() => {
    setPreviewLoading(true);
    return api
      .RenderAutomation({
        clusterAutomations: getKustomizations(),
      })
      .then(data => {
        setOpenPreview(true);
        setPRPreview(data);
      })
      .catch(err =>
        setNotifications([
          {
            message: { text: err.message },
            severity: 'error',
            display: 'bottom',
          },
        ]),
      )
      .finally(() => setPreviewLoading(false));
  }, [api, setOpenPreview, getKustomizations, setNotifications]);

  const token = getProviderToken(formData.provider);

  const { isAuthenticated, validateToken } = useIsAuthenticated(
    formData.provider,
    token,
  );

  const handleAddApplication = useCallback(() => {
    const payload: CreateAutomationsPullRequestRequest = {
      headBranch: formData.branchName,
      title: formData.pullRequestTitle,
      description: formData.pullRequestDescription,
      commitMessage: formData.commitMessage,
      clusterAutomations: getKustomizations(),
      repositoryUrl: getRepositoryUrl(formData.repo),
      baseBranch: formData.repo.obj.spec.ref.branch,
    };
    setLoading(true);
    return validateToken()
      .then(() =>
        api
          .CreateAutomationsPullRequest(payload, {
            headers: new Headers({
              'Git-Provider-Token': `token ${getProviderToken(
                formData.provider,
              )}`,
            }),
          })
          .then(response => {
            setPRPreview(null);
            history.push(Routes.Applications);
            setNotifications([
              {
                message: {
                  component: (
                    <Link href={response.webUrl} newTab>
                      PR created successfully, please review and merge the pull
                      request to apply the changes to the cluster.
                    </Link>
                  ),
                },
                severity: 'success',
              },
            ]);
          })
          .catch(error =>
            setNotifications([
              {
                message: { text: error.message },
                severity: 'error',
                display: 'bottom',
              },
            ]),
          )
          .finally(() => setLoading(false)),
      )
      .catch(() => {
        removeToken(formData.provider);
        setNotifications([expiredTokenNotification]);
      })
      .finally(() => setLoading(false));
  }, [
    api,
    formData,
    history,
    getKustomizations,
    setNotifications,
    validateToken,
  ]);

  const [submitType, setSubmitType] = useState<string>('');

  return useMemo(() => {
    return (
      <Page
        path={[
          {
            label: 'Applications',
            url: Routes.Applications,
          },
          { label: 'Add new application' },
        ]}
      >
        <CallbackStateContextProvider
          callbackState={{
            page: authRedirectPage as PageRoute,
            state: {
              formData,
              updatedProfiles,
            },
          }}
        >
          <NotificationsWrapper>
            <FormWrapper
              noValidate
              onSubmit={event =>
                validateFormData(
                  event,
                  submitType === 'PR Preview'
                    ? handlePRPreview
                    : handleAddApplication,
                  setFormError,
                  setSubmitType,
                )
              }
            >
              {formData.clusterAutomations.map(
                (
                  automation: FormData['clusterAutomations'][0],
                  index: number,
                ) => {
                  return (
                    <AppFields
                      context="app"
                      key={index}
                      index={index}
                      formData={formData}
                      setFormData={setFormData}
                      allowSelectCluster
                      clusterName={clusterName}
                      formError={formError}
                    />
                  );
                },
              )}
              {openPreview && prPreview ? (
                <Preview
                  context="app"
                  openPreview={openPreview}
                  setOpenPreview={setOpenPreview}
                  prPreview={prPreview}
                  sourceType={formData.source_type}
                />
              ) : null}
              <Box className="selected-source">
                {formData.source_url && 'Selected source: '}
                {optionUrl(formData.source_url, formData.source_branch)}
              </Box>
              {formData.source_type === 'HelmRepository' ? (
                <Profiles
                  cluster={{
                    name: formData.clusterAutomations[0].cluster_name,
                    namespace: formData.clusterAutomations[0].cluster_namespace,
                  }}
                  // Temp fix to hide layers when using profiles in Add App until we update the BE
                  context="app"
                  isLoading={profilesIsLoading}
                  updatedProfiles={updatedProfiles}
                  setUpdatedProfiles={setUpdatedProfiles}
                  helmRepo={helmRepo}
                />
              ) : null}
              {previewLoading ? (
                <LoadingPage className="preview-loading" />
              ) : (
                <Flex end className="preview-cta">
                  <Button
                    type="submit"
                    onClick={() => setSubmitType('PR Preview')}
                  >
                    PREVIEW PR
                  </Button>
                </Flex>
              )}
              <GitOps
                formData={formData}
                setFormData={setFormData}
                showAuthDialog={showAuthDialog}
                setShowAuthDialog={setShowAuthDialog}
                formError={formError}
                enableGitRepoSelection={true}
              />
              {loading ? (
                <LoadingPage className="create-loading" />
              ) : (
                <Flex end className="create-cta">
                  <Button
                    type="submit"
                    onClick={() => setSubmitType('Create app')}
                    disabled={!isAuthenticated}
                  >
                    CREATE PULL REQUEST
                  </Button>
                </Flex>
              )}
            </FormWrapper>
          </NotificationsWrapper>
        </CallbackStateContextProvider>
      </Page>
    );
  }, [
    authRedirectPage,
    formData,
    handleAddApplication,
    loading,
    profilesIsLoading,
    updatedProfiles,
    setUpdatedProfiles,
    showAuthDialog,
    prPreview,
    openPreview,
    handlePRPreview,
    previewLoading,
    clusterName,
    helmRepo,
    formError,
    submitType,
    isAuthenticated,
  ]);
};

export default ({ ...rest }) => <AddApplication {...rest} />;