import { FC } from 'react';
import { ContentWrapper } from '../Layout/ContentWrapper';
import { PageTemplate } from '../Layout/PageTemplate';
import { SectionHeader } from '../Layout/SectionHeader';
import { useApplicationsCount, useSourcesCount } from './utils';
import {
  OCIRepositoryDetail,
  Kind,
  useGetObject,
} from '@weaveworks/weave-gitops';
import { OCIRepository } from '@weaveworks/weave-gitops/ui/lib/objects';

type Props = {
  name: string;
  namespace: string;
  clusterName: string;
};

const WGApplicationsOCIRepository: FC<Props> = props => {
  const { name, namespace, clusterName } = props;
  const applicationsCount = useApplicationsCount();
  const sourcesCount = useSourcesCount();
  const {
    data: ociRepository,
    isLoading,
    error,
  } = useGetObject<OCIRepository>(
    name,
    namespace,
    Kind.OCIRepository,
    clusterName,
  );

  return (
    <PageTemplate documentTitle="WeGO · Git Repository">
      <SectionHeader
        path={[
          {
            label: 'Applications',
            url: '/applications',
            count: applicationsCount,
          },
          {
            label: 'Sources',
            url: '/sources',
            count: sourcesCount,
          },
          {
            label: `${props.name}`,
          },
        ]}
      />
      <ContentWrapper
        loading={isLoading}
        errors={
          error ? [{ clusterName, namespace, message: error?.message }] : []
        }
      >
        <OCIRepositoryDetail ociRepository={ociRepository} {...props} />
      </ContentWrapper>
    </PageTemplate>
  );
};

export default WGApplicationsOCIRepository;