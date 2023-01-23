import { useState } from 'react'
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import axios from "axios";

import './App.css'

const queryClient = new QueryClient();

function Example(props) {
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
      <Example api="http://localhost:8080/"/>
      <Example api="http://localhost:3000/"/>
    </QueryClientProvider>
  );
}

export default App
