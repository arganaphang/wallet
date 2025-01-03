import QueryProvider from "./components/query-provider"
import TransactionPage from "./components/transaction/page"

const App: React.FC = () => {
  return (
    <QueryProvider>
      <TransactionPage />
    </QueryProvider>
  )
}

export default App
