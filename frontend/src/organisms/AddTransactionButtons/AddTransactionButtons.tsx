import React from 'react';
import transferIcon from '../../assets/img/Transfer.png';
import topupIcon from '../../assets/img/Top-Up.png';
import { Button } from './AddTransactionButtons.Styles';

type AddTransactionProps = {
  onAddTransferClick: () => void;
  onAddTopupClick: () => void;
};

const AddTransaction = ({
  onAddTransferClick,
  onAddTopupClick,
}: AddTransactionProps): React.ReactElement => {
  return (
    <>
      <Button onClick={() => onAddTransferClick()}>
        <img src={transferIcon} alt=""></img>

        <p>Transfer +</p>
      </Button>

      <Button onClick={() => onAddTopupClick()}>
        <img src={topupIcon} alt=""></img>

        <p>Top Up +</p>
      </Button>
    </>
  );
};

export default AddTransaction;
