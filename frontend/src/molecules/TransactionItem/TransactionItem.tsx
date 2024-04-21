import React from 'react';
import { ITransaction } from '../../interfaces/Transaction';
import * as S from './TransactionItem.Styles';

type transactionItemProps = {
  transaction: ITransaction;
};

export const TransactionItem = ({
  transaction,
}: transactionItemProps): React.ReactElement => {
  const months = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
  ];

  const formatter = Intl.NumberFormat('en-US');

  function formatDate(dateStr: string): string {
    const [year, month, date] = dateStr.split('-');

    return `${date} ${months[+month.replace('0', '') - 1]} ${year}`;
  }

  return (
    <S.TransactionItemRow $transactionType={transaction.type}>
      <th>{formatDate(transaction.date)}</th>
      <th>{transaction.description === '' ? '-' : transaction.description}</th>
      <th>
        {transaction.description.includes('Top Up')
          ? '-'
          : transaction.type === 'income'
          ? transaction.sender_name
          : transaction.recipient_name}
      </th>
      <th>
        {transaction.type === 'income' ? '+' : '-'} IDR{' '}
        {formatter.format(+transaction.amount)}
      </th>
    </S.TransactionItemRow>
  );
};
