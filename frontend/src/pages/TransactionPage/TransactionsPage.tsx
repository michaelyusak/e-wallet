import React, { useEffect, useState } from 'react';

import * as S from './TransactionsPage.Styles';
import TransactionList from '../../organisms/TransactionList/TransactionList';
import {
  ITransaction,
  ITransactionResponse,
} from '../../interfaces/Transaction';
import SelectorGroup from '../../organisms/SelectorGroup/SelectorGroup';
import { ISelector } from '../../interfaces/Selector';
import PaginationButtons from '../../molecules/PaginationButton/PaginationButtons';
import AddTransaction from '../../organisms/AddTransactionButtons/AddTransactionButtons';
import TransactionModal from '../../organisms/TransactionModal/TransactionModal';
import { IFormField } from '../../interfaces/FormField';
import useFetch from '../../hooks/UseFetch';
import { useAppDispatch, useAppSelector } from '../../redux/ReduxHooks';
import { fetchUser } from '../../features/UserData/userSlice';
import LoadingModal from '../../molecules/LoadingModal/LoadingModal';

const Transactions = (): React.ReactElement => {
  const [page, setPage] = useState<number>(1);

  function handleSetPage(diff: number) {
    const maxPage: number = Math.ceil(
      (transactionData.data?.data.total_item || 0) / 8,
    );

    if (page === 1 && diff < 0) {
      return;
    }

    if (page === maxPage && diff > 0) {
      return;
    }

    setPage(page + diff);
  }

  const [sortBy, setSortBy] = useState<string>('date');
  const [sort, setSort] = useState<string>('desc');

  function handleSetSort(sort: string) {
    const [sortByParam, sortParam] = sort.split(' - ');

    setSortBy(sortByParam.toLowerCase());
    setSort(sortParam.toLowerCase());
  }

  const [search, setSearch] = useState<string>('All');

  const urlGetTransactions = `http://localhost:8080/transactions?search=${
    search === 'All' ? '' : search === 'Transfer' ? 'not Top Up' : search
  }&sortBy=${sortBy}&sort=${sort}&limit=8&page=${page}&from=&until=`;

  const [transactions, setTransactions] = useState<ITransaction[]>([]);

  const transactionData = useFetch<ITransactionResponse>(urlGetTransactions);

  const userData = useAppSelector((state) => state.user);

  const [isAddTransfer, setAddTransfer] = useState<boolean>(false);

  function handleSetAddTransfer() {
    setAddTransfer(!isAddTransfer);
  }

  const [isAddTopup, setAddTopup] = useState<boolean>(false);

  function handleSetAddTopup() {
    setAddTopup(!isAddTopup);
  }

  const dispatch = useAppDispatch();

  useEffect(() => {
    setPage(1);
    return () => {};
  }, [sortBy, sort, search]);

  useEffect(() => {
    dispatch(fetchUser());
    return () => {};
  }, [isAddTopup, isAddTransfer]);

  useEffect(() => {
    setTransactions(transactionData.data?.data.transactions ?? []);
    return () => {};
  }, [
    isAddTopup,
    isAddTransfer,
    transactionData.data?.data.transactions,
    urlGetTransactions,
  ]);

  const TransferFormFields: IFormField[] = [
    {
      name: 'destinationAccountNumber',
      placeholder: 'Enter destination account number',
      isRequired: true,
      type: 'text',
    },
    {
      type: 'currency',
      name: 'amount',
      label: 'Transfer',
      placeholder: 'Enter amount here',
      balance: userData.user.wallet.balance,
      isRequired: true,
    },
    {
      type: 'text',
      name: 'transferDescription',
      placeholder: 'Enter description',
    },
    {
      type: 'button',
      name: 'submit',
    },
  ];

  async function handleAddTransfer(inputValues: {
    [key: string]: { value: string; isError: boolean };
  }) {
    const to: string = inputValues['destinationAccountNumber'].value;
    const amount: number = +inputValues['amount'].value;
    const description: string = inputValues['transferDescription']?.value;

    const url = 'http://localhost:8080/transactions/transfer';
    const options = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: localStorage.getItem('token') || '',
      },
      body: JSON.stringify({
        to: to,
        amount: amount,
        description: description || '',
      }),
    };

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      if (!response.ok) {
        throw new Error(`Transfer failed ${responseData.message}`);
      }
    } catch (error) {
      console.log(error);
      throw error;
    }
  }

  const TopupFormFields: IFormField[] = [
    {
      type: 'currency',
      name: 'amount',
      label: 'Top Up',
      placeholder: 'Enter amount here',
      isRequired: true,
    },
    {
      type: 'button',
      name: 'submit',
    },
  ];

  const TopupDropDown: ISelector = {
    title: undefined,
    initialValue: 'Choose source of funds',
    menus: ['Credit Card', 'Cash', 'Rewards'],
    setVal: (val: string) => handleSetSourceOfFund(val),
  };

  const [sourceOfFund, setSourceOfFund] = useState<number>(0);

  function handleSetSourceOfFund(val: string) {
    if (val === 'Credit Card') {
      setSourceOfFund(3);
    } else if (val === 'Cash') {
      setSourceOfFund(4);
    } else if (val === 'Rewards') {
      setSourceOfFund(5);
    }
  }

  async function handleAddTopup(inputValues: {
    [key: string]: { value: string; isError: boolean };
  }) {
    const amount: number = +inputValues['amount'].value;
    const source_of_funds: number = sourceOfFund;

    const url = 'http://localhost:8080/transactions/topup';
    const options = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: localStorage.getItem('token') || '',
      },
      body: JSON.stringify({
        amount: amount,
        source_of_funds: source_of_funds,
      }),
    };

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      if (!response.ok) {
        throw new Error(`Topup failed ${responseData.message}`);
      }
    } catch (error) {
      console.log(error);
      throw error;
    }
  }

  const selectors: ISelector[] = [
    {
      title: 'Type',
      menus: ['All', 'Transfer', 'Top Up'],
      initialValue: 'All',
      setVal: (val: string) => setSearch(val),
    },
    {
      title: 'Sort',
      menus: ['Amount - Asc', 'Amount - Desc', 'Date - Asc', 'Date - Desc'],
      initialValue: 'Date - Desc',
      setVal: (val: string) => handleSetSort(val),
    },
  ];

  const formatter = new Intl.NumberFormat('en-US');

  const [isShowBalance, setShowBalance] = useState<boolean>(false);

  function handleSetShowBalance() {
    setShowBalance(!isShowBalance);
  }

  if (userData.error) return <S.Status>Error: {userData.error}</S.Status>;

  if (transactionData.error)
    return <S.Status>Error: {transactionData.error?.message}</S.Status>;

  return (
    <>
      {(userData.isLoading || transactionData.isLoading) && (
        <LoadingModal></LoadingModal>
      )}

      <S.FullPage>
        <S.TransactionsContent>
          <S.Hero>
            <button onClick={() => handleSetShowBalance()}>
              <h1>
                IDR{' '}
                {isShowBalance
                  ? formatter.format(+userData.user.wallet.balance)
                  : '*'.repeat(userData.user.wallet.balance.length || 1)}
              </h1>
            </button>

            <p>Total balance from account {userData.user.wallet.number || 1000000000000}</p>
          </S.Hero>

          <S.Transactions>
            <S.Selectors>
              <SelectorGroup selectors={selectors}></SelectorGroup>
            </S.Selectors>

            <S.TransactionListContainer>
              <TransactionList transactions={transactions}></TransactionList>
            </S.TransactionListContainer>

            <PaginationButtons
              setPage={(diff) => handleSetPage(diff)}
            ></PaginationButtons>
          </S.Transactions>

          <S.AddTransactionButtonsContainer>
            <AddTransaction
              onAddTransferClick={() => handleSetAddTransfer()}
              onAddTopupClick={() => handleSetAddTopup()}
            ></AddTransaction>
          </S.AddTransactionButtonsContainer>
        </S.TransactionsContent>

        {isAddTransfer && (
          <TransactionModal
            title="Transfer"
            formFields={TransferFormFields}
            onSubmit={(inputValues) => handleAddTransfer(inputValues)}
            closeModal={() => handleSetAddTransfer()}
          ></TransactionModal>
        )}

        {isAddTopup && (
          <TransactionModal
            title="Top Up"
            formFields={TopupFormFields}
            dropDown={TopupDropDown}
            onSubmit={(inputValues) => handleAddTopup(inputValues)}
            closeModal={() => handleSetAddTopup()}
          ></TransactionModal>
        )}
      </S.FullPage>
    </>
  );
};

export default Transactions;
