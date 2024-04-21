import React, { useContext, useState } from 'react';
import * as S from './TransactionModal.Styles';
import Form from '../Form/Form';
import { IFormField } from '../../interfaces/FormField';
import { ISelector } from '../../interfaces/Selector';
import TopupDropdown from '../../molecules/TopupDropdown/TopupDropdown';
import checkIcon from '../../assets/img/check-icon.png';
import { ToastContext } from '../../contexts/ToastData';

type TransactionModalProps = {
  title: string;
  formFields: IFormField[];
  dropDown?: ISelector;
  closeModal: () => void;
  onSubmit: (inputValues: {
    [key: string]: { value: string; isError: boolean };
  }) => Promise<void>;
};

const TransactionModal = ({
  title,
  formFields,
  dropDown,
  onSubmit,
  closeModal,
}: TransactionModalProps): React.ReactElement => {
  const formatter = new Intl.NumberFormat('en-US');
  const { setToast } = useContext(ToastContext);

  function handleSetToast(isSuccess: boolean, message: string) {
    setToast((prevToast) => ({
      ...prevToast,
      isSuccess: isSuccess,
      message: message,
      isVisible: true,
      marginLeft: '41%',
      marginTop: '5%',
    }));

    setTimeout(
      () =>
        setToast((prevToast) => ({
          ...prevToast,
          isSuccess: undefined,
          message: '',
          isVisible: false,
          marginLeft: '',
          marginTop: '',
        })),
      5000,
    );
  }

  async function handleOnSubmit(inputValues: {
    [key: string]: { value: string; isError: boolean };
  }) {
    try {
      await onSubmit(inputValues);

      handleSetToast(true, `${title} Success`);

      setComplete(true);
      setTransactionAmount(formatter.format(+inputValues['amount'].value));
    } catch (error) {
      console.log(error);

      handleSetToast(false, `${title} Failed`);
    }
  }

  const [transactionAmount, setTransactionAmount] = useState<string>('');

  const [isComplete, setComplete] = useState<boolean>(false);

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

  function getDate(): string {
    const d = new Date();

    return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
  }

  return (
    <S.ModalContainer>
      <S.ModalUnderLay onClick={() => closeModal()}></S.ModalUnderLay>

      {!isComplete ? (
        <S.TransactionCard>
          <h1>{title}</h1>
          <S.TransactionForm>
            {dropDown && <TopupDropdown dropdown={dropDown}></TopupDropdown>}

            <Form
              formFields={formFields}
              gap="10px"
              onSubmit={(inputValues) => handleOnSubmit(inputValues)}
            ></Form>
          </S.TransactionForm>
        </S.TransactionCard>
      ) : (
        <S.TransactionCard $gap="30px">
          <S.ImgContainer>
            <img src={checkIcon} alt=""></img>
          </S.ImgContainer>

          <h2>{title} Success</h2>

          <h3>IDR {transactionAmount}</h3>

          <p>{`${getDate()}`}</p>

          <S.Button onClick={() => closeModal()}>Close</S.Button>
        </S.TransactionCard>
      )}
    </S.ModalContainer>
  );
};

export default TransactionModal;
