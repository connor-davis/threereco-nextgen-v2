import { getApiAuthenticationCheckOptions } from '@/api-client/@tanstack/react-query.gen';
import { useQuery } from '@tanstack/react-query';
import { type ReactNode, createContext, useContext } from 'react';

import type { User } from '@/api-client';
import { apiClient } from '@/lib/utils';

const AuthenticationContext = createContext<{
  user?: User;
  isLoading: boolean;
  isError: boolean;
  refetch: () => Promise<void>;
}>({
  user: undefined,
  isLoading: false,
  isError: false,
  refetch: async () => {},
});

export const AuthenticationProvider = ({
  children,
}: {
  children: ReactNode;
}) => {
  const {
    data: userData,
    isLoading: isLoadingUser,
    isError: isUserError,
    refetch: refetchUser,
  } = useQuery({
    ...getApiAuthenticationCheckOptions({
      client: apiClient,
    }),
    retry: 0,
    refetchOnWindowFocus: true,
    refetchOnMount: true,
    refetchOnReconnect: true,
  });

  return (
    <AuthenticationContext.Provider
      value={{
        user: userData?.item as User | undefined,
        isLoading: isLoadingUser,
        isError: isUserError,
        refetch: async () => {
          await refetchUser();
        },
      }}
    >
      {children}
    </AuthenticationContext.Provider>
  );
};

export const useAuthentication = () => useContext(AuthenticationContext);
