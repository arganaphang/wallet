import { Transaction, TransactionType } from "@/types/transaction";
import { cn, CurrencyHelper } from "@/utils";

type Props = {
  transaction: Transaction;
}

const TransactionItem: React.FC<Props> = ({ transaction }) => {
  return (
    <div className="flex bg-[#c4a9ff12] px-4 py-2 rounded-xl">
      <div className="flex-1 flex items-center gap-2">
        <div className={cn("w-8 h-8 rounded-xl", {
          "bg-green-200": transaction.type === TransactionType.INCOME,
          "bg-red-200": transaction.type === TransactionType.OUTCOME
        })}></div>
        <p className="text-sm">{transaction.name}</p>
      </div>
      <div className="flex flex-col gap-2 items-end">
        <p className=" text-xs font-bold">{CurrencyHelper.toRupiah(transaction.amount)}</p>
        <small className="border rounded-full text-center px-4 py-0.5">{transaction.category}</small>
      </div>
    </div>
  );
}

export default TransactionItem;
