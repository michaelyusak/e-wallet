import styled from 'styled-components';

export const InputSignInPassword = styled.div<{ $isError?: boolean }>`
  div {
    display: flex;
    justify-content: space-between;
    align-items: center;

    border-radius: 8px;
    background-color: ${(props) => (props.$isError ? '#FEEFE7' : '#f0efff')};
    border: ${(props) => props.$isError && 'solid 1px #F60707'};
    height: 62px;
    padding: 20px 26px;

    input {
      border: none;
      outline: none;
      width: 100%;
      background-color: transparent;
      color: black;
      line-height: 22.5px;
      font-size: 15px;
      font-weight: 400;
      letter-spacing: 12px;
    }

    input::placeholder {
      letter-spacing: normal;
      color: ${(props) => (props.$isError ? '#F60707' : '#a7a3ff')};
    }

    button {
      height: 17px;
      border: none;
      outline: none;
      background-color: transparent;

      img {
        width: 17px;
        height: 17px;
        cursor: pointer;
      }
    }
  }

  p {
    text-align: right;
    margin-top: 17px;

    color: #b0b0b0;
    line-height: 19.5px;
    font-size: 13px;
    font-weight: 400;

    cursor: pointer;
  }
`;

export const InputField = styled.div<{
  $isError?: boolean;
  $fontColor?: string;
}>`
  border-radius: 8px;
  background-color: ${(props) => (props.$isError ? '#FEEFE7' : '#f0efff')};
  border: ${(props) => props.$isError && 'solid 1px #F60707'};
  height: 62px;
  display: flex;
  align-items: center;
  padding: 20px 26px;

  input {
    border: none;
    outline: none;
    width: 100%;
    background-color: transparent !important;
    color: ${(props) => props.$fontColor || 'black'};
    line-height: 22.5px;
    font-size: 15px;
    font-weight: 400;
  }

  input::placeholder {
    color: ${(props) => (props.$isError ? '#F60707' : '#a7a3ff')};
  }
`;

export const InputPassword = styled.div<{ $isError?: boolean }>`
  border-radius: 8px;
  background-color: ${(props) => (props.$isError ? '#FEEFE7' : '#f0efff')};
  border: ${(props) => props.$isError && 'solid 1px #F60707'};
  height: 62px;
  display: flex;
  align-items: center;
  padding: 20px 26px;

  input {
    border: none;
    outline: none;
    width: 100%;
    background-color: transparent;
    color: black;
    line-height: 22.5px;
    font-size: 15px;
    font-weight: 400;
    letter-spacing: 12px;
  }

  button {
    height: 17px;
    border: none;
    outline: none;
    background-color: transparent;

    img {
      width: 17px;
      height: 17px;
      cursor: pointer;
    }
  }

  input::placeholder {
    color: ${(props) => (props.$isError ? '#F60707' : '#a7a3ff')};
    letter-spacing: normal;
  }
`;

export const InputCurrency = styled.div<{
  $Error?: 'invalidAmount' | 'insufficientBalance';
}>`
  display: flex;
  flex-direction: column;

  > div {
    display: flex;
    height: 59px;
    align-items: center;
    background-color: #f0efff;
    border-radius: 8px;
    padding: 14px 20px 14px 0px;
    gap: 20px;

    > div {
      background-color: #4d47c3;
      height: 59px;
      border-radius: 8px 0px 0px 8px;
      padding: 11px 15px;
      display: flex;
      align-items: center;

      p {
        line-height: 22.5px;
        font-size: 15px;
        font-weight: 400;
        color: white;
      }
    }

    input {
      border: none;
      outline: none;
      width: 100%;
      background-color: transparent;
      color: black;
      line-height: 22.5px;
      font-size: 15px;
      font-weight: 400;
    }

    input::placeholder {
      color: #a7a3ff;
    }
  }
`;

export const Balance = styled.p`
  margin-top: 10px;
  line-height: 20px;
  font-size: 14px;
  font-weight: 400;
  color: #95999e;
`;

export const ErrorStatement = styled.p`
  margin-top: 10px;
  line-height: 20px;
  font-size: 14px;
  font-weight: 700;
  color: #f60707;
`;
