import React from 'react';
import * as S from './RecentTransactionItem.Styles';
import incomeIcon from '../../assets/img/income.png';
import expenseIcon from '../../assets/img/expense.png';
import { ITransaction } from '../../interfaces/Transaction';

type RecentTransactionCardProps = {
  recentTransaction: ITransaction;
};

const RecentTransactionCard = ({
  recentTransaction,
}: RecentTransactionCardProps): React.ReactElement => {
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
    <>
      <S.RecentTransactionCard $type={recentTransaction.type}>
        <img
          src={recentTransaction.type === 'income' ? incomeIcon : recentTransaction.type === 'expense' ? expenseIcon : undefined}
          alt=""
        ></img>

        <p>{recentTransaction.description || '-'}</p>

        <p>{formatDate(recentTransaction.date)}</p>

        <p>IDR {formatter.format(+recentTransaction.amount)}</p>
      </S.RecentTransactionCard>
    </>
  );
};

export default RecentTransactionCard;
