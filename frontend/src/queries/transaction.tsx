import { AddTransactionRequest, Transaction, UpdateTransactionRequest } from "@/types/transaction";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

const BASE_URL = "http://0.0.0.0:8000/api/v1/transactions";

const getAllTransactions = async () => {
  const response = await fetch(`${BASE_URL}`);
  const result = await response.json();
  return result.data as Transaction[];
}

export const useGetAllTransactions = () => useQuery({
  queryKey: ['transactions'],
  queryFn: getAllTransactions,
});


const getTransactionByID = async (id: string) => {
  const response = await fetch(`${BASE_URL}/${id}`);
  const result = await response.json();
  return result.data as Transaction;
}

export const useGetTransactionByID = (id: string) => useQuery({
  queryKey: ['transactions', { id: id }],
  queryFn: () => getTransactionByID(id),
});

const addTransaction = async (data: AddTransactionRequest) => {
  await fetch(`${BASE_URL}`, {
    method: 'POST',
    body: JSON.stringify(data)
  });
  return;
}

export const useMutationAddTransaction = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: addTransaction,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['transactions'] })
    }
  });
}

const updateTransactionByID = async (data: UpdateTransactionRequest) => {
  await fetch(`${BASE_URL}/${data.id}`, {
    method: 'PUT',
    body: JSON.stringify(data)
  });
  return;
}

export const useMutationUpdateTransactionByID = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: updateTransactionByID,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['transactions'] })
    }
  });
}

const deleteTransactionByID = async (id: string) => {
  await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
  });
  return;
}

export const useMutationDeleteTransactionByID = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: deleteTransactionByID,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['transactions'] })
    }
  });
}
