export interface IRecentTransaction {
  type: 'income' | 'expense';
  description: string;
  date: string;
  amount: string;
}
