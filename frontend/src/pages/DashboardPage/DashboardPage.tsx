import React, { useEffect, useState } from 'react';
import useFetch from '../../hooks/UseFetch';
import { useAppDispatch, useAppSelector } from '../../redux/ReduxHooks';

import {
  ITransaction,
  ITransactionResponse,
} from '../../interfaces/Transaction';

import RecentTransactionList from '../../organisms/RecentTransactionList/RecentTransactionList';
import WalletInformation from '../../organisms/WalletInformation/WalletInformation';

import { fetchUser } from '../../features/UserData/userSlice';

import * as S from './DashboardPage.Styles';
import LoadingModal from '../../molecules/LoadingModal/LoadingModal';

const Dashboard = (): React.ReactElement => {
  function getDate(): [string, string] {
    const d = new Date();

    const dateLastWeek = `${d.getFullYear()}-${('0' + (d.getMonth() + 1)).slice(
      -2,
    )}-${('0' + (d.getDate() - 7)).slice(-2)}`;
    const dateNow = `${d.getFullYear()}-${('0' + (d.getMonth() + 1)).slice(
      -2,
    )}-${('0' + d.getDate()).slice(-2)}`;

    return [dateLastWeek, dateNow];
  }
  const [dateLastWeek, dateNow] = getDate();

  const urlGetLastWeekTransactions = `http://localhost:8080/transactions?search=&sortBy=&sort=desc&limit=3&page=1&from=${dateLastWeek}&until=${dateNow}`;

  const [lastWeekTransactions, setLastWeekTransaction] = useState<
    ITransaction[]
  >([]);

  const lastWeekTransactionData = useFetch<ITransactionResponse>(
    urlGetLastWeekTransactions,
  );

  const userData = useAppSelector((state) => state.user);

  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(fetchUser());
    return () => {};
  }, []);

  useEffect(() => {
    setLastWeekTransaction(
      lastWeekTransactionData.data?.data.transactions ?? [],
    );
    return () => {};
  }, [lastWeekTransactionData.data?.data]);

  if (userData.error) return <S.Status>Error: {userData.error}</S.Status>;

  if (lastWeekTransactionData.error)
    return <S.Status>Error: {lastWeekTransactionData.error?.message}</S.Status>;

  return (
    <S.FullPage>
      {userData.isLoading && <LoadingModal></LoadingModal>}
      <S.DashboardContent>
        <h2>Hello, {userData.user?.name}!</h2>

        <S.WalletInformation>
          <WalletInformation wallet={userData.user?.wallet}></WalletInformation>
        </S.WalletInformation>

        <S.RecentTransactions>
          <S.RecentTransactionsHeader>
            <h2>Recent Transactions</h2>

            <p>This Week</p>
          </S.RecentTransactionsHeader>

          <RecentTransactionList
            recentTransactions={lastWeekTransactions}
          ></RecentTransactionList>
        </S.RecentTransactions>
      </S.DashboardContent>
    </S.FullPage>
  );
};

export default Dashboard;
