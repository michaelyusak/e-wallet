import styled from 'styled-components';

export const DropdownContainer = styled.div`
  > button {
    display: flex;
    align-items: center;
    width: 100%;
    justify-content: space-between;
    height: 59px;
    padding: 10px 20px;
    border-radius: 8px;
    background-color: #f0efff;
    border: none;
    outline: none;
    cursor: pointer;

    p {
      color: #a7a3ff;
      font-size: 15px;
      font-weight: 400;
      line-height: 22.5px;
    }
  }

  div:nth-child(2) {
    width: 84%;
    display: flex;
    flex-direction: column;
    position: absolute;
    margin-top: 10px;
    z-index: 2;
    background-color: #f0efff;
    padding: 10px 20px;
    border-radius: 8px;
    box-shadow: 0px 4px 4px 0px #00000040;
    gap: 10px;

    > button {
      background-color: transparent;
      border: none;
      outline: none;
      cursor: pointer;
      height: 30px;
      line-height: 22.5px;
      font-size: 15px;
      font-weight: 400;
      color: #a7a3ff;
      text-align: left;
    }
  }
`;

export const Img = styled.img<{ $direction: 'up' | 'down' }>`
  rotate: ${(props) => props.$direction === 'up' && '0deg'};
  rotate: ${(props) => props.$direction === 'down' && '180deg'};
`;
