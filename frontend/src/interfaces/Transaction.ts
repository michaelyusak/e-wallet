import { IResponse } from './APIResponse';

export interface ITransaction {
  id: string;
  type: 'income' | 'expense';
  from: string;
  sender_name: string;
  to: string;
  recipient_name: string;
  amount: string;
  source_of_funds: string;
  description: string;
  date: string;
}

export interface ITransactionList {
  total_item: number;
  transactions: ITransaction[];
}

export interface ITransactionResponse extends IResponse {
  data: ITransactionList;
}
