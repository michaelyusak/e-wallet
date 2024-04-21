import React, { useState } from 'react';
import { ISelector } from '../../interfaces/Selector';
import * as S from './Selector.Styles';

type selectorProps = {
  selector: ISelector;
};

const Selector = ({ selector }: selectorProps): React.ReactElement => {
  const [val, setVal] = useState<string>(selector.initialValue);

  const [isDropped, setDropped] = useState<boolean>(false);

  function handleSelectMenu(menu: string) {
    setVal(menu);

    selector.setVal(menu);

    setDropped(false);
  }

  return (
    <S.SelectorContainer>
      {selector.title && <p>{selector.title}</p>}

      <div>
        <button onClick={() => setDropped(!isDropped)}>{val}</button>

        {isDropped && (
          <S.DropdownMenu>
            {selector.menus.map((menu) => (
              <button key={menu} onClick={() => handleSelectMenu(menu)}>
                {menu}
              </button>
            ))}
          </S.DropdownMenu>
        )}
      </div>
    </S.SelectorContainer>
  );
};

export default Selector;
