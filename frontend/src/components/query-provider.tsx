import * as React from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

interface Props extends React.PropsWithChildren<{}> { }

const QueryProvider: React.FC<Props> = (props: Props) => {
  const [client, _] = React.useState(new QueryClient());
  return (
    <QueryClientProvider client={client}>
      {props.children}
      <ReactQueryDevtools initialIsOpen={false} buttonPosition="bottom-left" />
    </QueryClientProvider>
  );
}

export default QueryProvider;
