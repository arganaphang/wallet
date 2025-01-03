import { useQuery } from "@tanstack/react-query";
import { Category } from "@/types/category";

const BASE_URL = "http://0.0.0.0:8000/api/v1/categories";

const getAllCategories = async () => {
  const response = await fetch(`${BASE_URL}`);
  const result = await response.json();
  return result.data as Category[];
}

export const useGetAllCategories = () => useQuery({
  queryKey: ['categories'],
  queryFn: getAllCategories
});
