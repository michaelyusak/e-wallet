import React from 'react';
import { ITransaction } from '../../interfaces/Transaction';
import { TransactionItem } from '../../molecules/TransactionItem/TransactionItem';
import * as S from './TransactionList.Styles';

type transactionListProps = {
  transactions: ITransaction[];
};

const TransactionList = ({
  transactions,
}: transactionListProps): React.ReactElement => {
  return (
    <S.TransactionListTable>
      <thead>
        <tr>
          <th>DATE</th>
          <th>DESCRIPTION</th>
          <th>TO / FROM</th>
          <th>AMOUNT</th>
        </tr>
      </thead>

      <tbody>
        {transactions.map((transaction) => (
          <TransactionItem
            transaction={transaction}
            key={transaction.id}
          ></TransactionItem>
        ))}
      </tbody>
    </S.TransactionListTable>
  );
};

export default TransactionList;
