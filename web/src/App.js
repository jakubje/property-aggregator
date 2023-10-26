import "./App.css";
import { QueryClient, QueryClientProvider } from "react-query";
import Properties from "./components/Properties";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        <Properties />
      </div>
    </QueryClientProvider>
  );
}

export default App;
