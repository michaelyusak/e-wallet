import styled from 'styled-components';

export const ModalContainer = styled.div`
  position: absolute;
  top: 0;
  z-index: 5;
  width: 100%;
  height: 100%;
`;

export const ModalUnderLay = styled.div`
  background-color: #454545;
  opacity: 90%;
  width: 100%;
  height: 100%;
`;

export const TransactionCard = styled.div<{ $gap?: string }>`
  position: absolute;
  z-index: 2;
  top: 15%;
  left: 35%;
  width: 30%;
  height: 70%;
  background-color: white;
  border-radius: 9px;
  padding: 90px 44px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  box-shadow: 0px 4px 4px 0px #00000040;
  gap: ${(props) => props.$gap || '0px'};

  h1 {
    line-height: 63px;
    font-size: 42px;
    font-weight: 600;
    color: #4d47c3;
    margin-bottom: 50px;
  }

  h2 {
    line-height: 42px;
    font-size: 28px;
    font-weight: 600;
    color: #4d47c3;
    text-align: center;
  }

  h3 {
    line-height: 48px;
    font-size: 32px;
    font-weight: 600;
    color: #282828;
    text-align: center;
  }

  > p {
    line-height: 24px;
    font-size: 16px;
    font-weight: 600;
    color: #95999e;
    text-align: center;
  }
`;

export const ImgContainer = styled.div`
  height: 127.5px;
  width: 127.5px;
  background-color: #33a720;
  border-radius: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0 auto;

  img {
    height: 100px;
    width: 100px;
  }
`;

export const Button = styled.button`
  margin: 0 auto;
  width: 100%;
  height: 59px;
  width: 369px;
  border-radius: 9px;
  background-color: #4d47c3;
  border: none;
  outline: none;
  box-shadow: 0px 1px 20px 0px #4d47c366;
  line-height: 24px;
  font-size: 16px;
  font-weight: 500;
  color: white;
  cursor: pointer;
`;

export const TransactionForm = styled.div`
  display: flex;
  flex-direction: column;
  gap: 10px;

  input[type='submit'] {
    margin-top: 10px;
  }
`;
