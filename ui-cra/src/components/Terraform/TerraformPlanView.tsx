import styled from 'styled-components';
import CodeView from '../CodeView';
import { Body, Message, Title } from '../Shared';
import { Link } from '@weaveworks/weave-gitops';

type Props = {
  plan?: string;
  error?: string;
};

function TerraformPlanView({ plan, error }: Props) {
  return (
    <>
      {plan && !error ? (
        <CodeView
          kind="Terraform"
          code={plan.trimStart() || ''}
          colorizeChanges
        />
      ) : (
        <Message>
          <Title>Terraform Plan</Title>
          <Body>No plan available.</Body>
          <Body>
            To enable the plan view, please set the field
            `spec.storeReadablePlan` to `human`.
          </Body>
          <Body>
            To learn more about planning Terraform resources,&nbsp;
            <Link
              href="https://docs.gitops.weave.works/docs/terraform/Using%20Terraform%20CR/plan-and-manually-apply-terraform-resources/"
              newTab
            >
              visit our documentation
            </Link>
          </Body>
        </Message>
      )}
    </>
  );
}

export default styled(TerraformPlanView).attrs({
  className: TerraformPlanView.name,
})``;