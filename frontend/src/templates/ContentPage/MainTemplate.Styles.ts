import styled from 'styled-components';

export const SeparatorLine = styled.div`
  position: absolute;
  top: 93px;
  border-bottom: #ede6e7 1px solid;
  width: 100%;
  height: 0px;

  @media (max-width: 710px) {
    margin-top: 85px;
  }
`;

export const HeaderContainer = styled.div`
  position: sticky;
  top: 0;
`;

export const FullPage = styled.section`
  height: 100vh;

  @media (max-width: 1050px) {
    height: fit-content;
  }
`;
