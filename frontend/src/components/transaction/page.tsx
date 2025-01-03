import { useGetAllTransactions } from "@/queries/transaction";
import React from "react";
import TransactionItem from "./transaction-item";

const TransactionPage: React.FC = () => {
  const transactions = useGetAllTransactions();

  return (
    <main className="px-4 py-6 flex flex-col gap-4">
      <h1 className="text-2xl">Transaction</h1>

      {!transactions.data && <div className="px-8 py-4 flex justify-center items-center bg-[#c4a9ff12]">
        <h1>Transaction Empty 🙏</h1>
      </div>}

      <div>
        {transactions.data?.map(trx => <TransactionItem key={trx.id} transaction={trx} />)}
      </div>
      <button className="fixed bottom-8 right-8 w-16 h-16 flex justify-center items-center rounded-xl bg-[#c4a9ff12]">+</button>
    </main>
  );
};

export default TransactionPage
