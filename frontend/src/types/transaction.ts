export const TransactionType = {
  INCOME: "income",
  OUTCOME: "outcome",
} as const;

export type Transaction = {
  id: string;
  name: string;
  amount: number;
  category: string;
  type: (typeof TransactionType)[keyof typeof TransactionType];
  created_at: Date;
};

export type AddTransactionRequest = {
  name: string;
  amount: string;
  category: string;
  type: (typeof TransactionType)[keyof typeof TransactionType];
}

export type UpdateTransactionRequest = {
  id: string;
  name: string;
  amount: string;
  category: string;
  type: (typeof TransactionType)[keyof typeof TransactionType];
}
