import styled, { keyframes } from 'styled-components';

const pop = keyframes`
    0% {
        transform: translateY(-30px);
        opacity: 0;
    }

    10% {
        transform: translateY(0);
        opacity: 1;
    }

    90% {
        transform: translateY(0px);
        opacity: 1;
    }

    100% {
        transform: translateY(-30px);
        opacity: 0;
    }
`;

export const Toast = styled.div<{
  $isSucces: boolean;
  $marginTop?: string;
  $marginLeft?: string;
}>`
  position: absolute;
  z-index: 10;
  padding: 5px 103px;
  background-color: ${(props) => (props.$isSucces ? '#EAFCEF' : '#FFDDCA')};
  border: solid 0.5px ${(props) => (props.$isSucces ? '#33A720' : '#F60707')};
  border-radius: 8px;
  margin-top: ${(props) => props.$marginTop || '7%'};
  margin-left: ${(props) => props.$marginLeft || '75%'};
  box-shadow: 0px 2px 2px 0px #00000040;
  animation: ${pop} 3s linear forwards;
  line-height: 24px;
  font-size: 16px;
  font-weight: 500;
  color: ${(props) => (props.$isSucces ? '#33A720' : '#F60707')};
`;
