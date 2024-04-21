import React from 'react';

import * as S from './RecentTransactionList.Styles';
import RecentTransactionCard from '../../molecules/RecentTransactionItem/RecentTransactionItem';
import { ITransaction } from '../../interfaces/Transaction';

type RecentTransactionProps = {
  recentTransactions: ITransaction[];
};

const RecentTransactionList = ({
  recentTransactions,
}: RecentTransactionProps): React.ReactElement => {
  return (
    <S.RecentTransactionList>
      {recentTransactions.length === 0 ? (
        <S.NoTransactionStatement>
          <h3>No recent transactions</h3>

          <p>Go to transactions page to add some!</p>
        </S.NoTransactionStatement>
      ) : (
        recentTransactions.map((recentTransaction) => (
          <>
            <RecentTransactionCard
              recentTransaction={recentTransaction}
            ></RecentTransactionCard>
          </>
        ))
      )}
    </S.RecentTransactionList>
  );
};

export default RecentTransactionList;
