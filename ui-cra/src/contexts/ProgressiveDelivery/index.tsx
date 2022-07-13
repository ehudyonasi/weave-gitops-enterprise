import {
  GetCanaryResponse,
  ListCanariesResponse,
  ProgressiveDeliveryService,
} from '@weaveworks/progressive-delivery';
import _ from 'lodash';
import React, { useContext } from 'react';
import { useQuery } from 'react-query';
import {
  ListEventsRequest,
  ListEventsResponse,
} from '../../cluster-services/cluster_services.pb';
import { EnterpriseClientContext } from '../EnterpriseClient';

interface Props {
  api: typeof ProgressiveDeliveryService;
  children: any;
}

export const ProgressiveDeliveryContext = React.createContext<
  typeof ProgressiveDeliveryService
>(null as any);

export const ProgressiveDeliveryProvider = ({ api, children }: Props) => (
  <ProgressiveDeliveryContext.Provider value={api}>
    {children}
  </ProgressiveDeliveryContext.Provider>
);

export const useProgressiveDelivery = () =>
  React.useContext(ProgressiveDeliveryContext);

const PD_QUERY_KEY = 'flagger';
const FLAGGER_STATUS_KEY = 'status';

export const useIsFlaggerAvailable = () => {
  const pd = useProgressiveDelivery();

  return useQuery<boolean, Error>([PD_QUERY_KEY, FLAGGER_STATUS_KEY], () => {
    return pd.IsFlaggerAvailable({}).then(res => {
      if (!res.clusters) {
        return false;
      }

      return _.includes(_.values(res.clusters), true);
    });
  });
};

const CANARIES_KEY = 'canaries';
export const useListCanaries = () => {
  const pd = useProgressiveDelivery();

  return useQuery<ListCanariesResponse, Error>(
    [PD_QUERY_KEY, CANARIES_KEY],
    () => pd.ListCanaries({}),
  );
};

type CanaryParams = {
  name: string;
  namespace: string;
  clusterName: string;
};

export const useGetCanaryDetails = (params: CanaryParams) => {
  const pd = useProgressiveDelivery();

  return useQuery<GetCanaryResponse, Error>(
    [PD_QUERY_KEY, CANARIES_KEY, params],
    () => pd.GetCanary(params),
  );
};

export const useCanariesCount = () => {
  const { data } = useListCanaries();

  if (!data) {
    return 0;
  }

  return data?.canaries?.length;
};

const EVENTS_QUERY_KEY = 'events';

export function useListEvents(req: ListEventsRequest) {
  const { api } = useContext(EnterpriseClientContext);

  return useQuery<ListEventsResponse, Error>([EVENTS_QUERY_KEY, req], () =>
    api.ListEvents(req),
  );
}