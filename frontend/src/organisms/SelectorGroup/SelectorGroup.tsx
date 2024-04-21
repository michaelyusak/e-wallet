import React from 'react';
import { ISelector } from '../../interfaces/Selector';
import Selector from '../../molecules/Selector/Selector';

type selectorGroupProps = {
  selectors: ISelector[];
};

const SelectorGroup = ({
  selectors,
}: selectorGroupProps): React.ReactElement => {
  return (
    <>
      {selectors.map((selector) => (
        <Selector selector={selector} key={selector.title}></Selector>
      ))}
    </>
  );
};

export default SelectorGroup;
