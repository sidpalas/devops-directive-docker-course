import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import axios from "axios";

import './App.css'

const queryClient = new QueryClient();

function CurrentTime(props) {
  const { isLoading, error, data, isFetching } = useQuery({
    queryKey: [props.api],
    queryFn: () =>
      axios
        .get(`${props.api}`)
        .then((res) => res.data),
  });

  if (isLoading) return `Loading ${props.api}... `;

  if (error) return "An error has occurred: " + error.message;

  return (
    <div className="App">
      <p>---</p>
      <p>API: {data.api}</p>
      <p>Time from DB: {data.now}</p>
      <div>{isFetching ? "Updating..." : ""}</div>
    </div>
  )
}

export function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <h1>Hey Team! 👋</h1>
      <CurrentTime api="/api/golang/"/>
      <CurrentTime api="/api/node/"/>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App
