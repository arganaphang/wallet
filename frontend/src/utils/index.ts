import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export const CurrencyHelper = {
  toRupiah: (value: number) => {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value);
  }
};

export const DateHelper = {
  formatDate: (input: string | number): string => {
    const date = new Date(input)
    return date.toLocaleDateString("en-US", {
      month: "long",
      day: "numeric",
      year: "numeric",
    })
  }
};

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

