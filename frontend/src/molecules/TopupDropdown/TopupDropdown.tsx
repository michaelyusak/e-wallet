import React, { useState } from 'react';
import * as S from './TopupDropdown.Styles';
import dropdownArrorIcon from '../../assets/img/dropdown-arrow.png';
import { ISelector } from '../../interfaces/Selector';

type TopupDropdownProps = {
  dropdown: ISelector;
};

const TopupDropdown = ({
  dropdown,
}: TopupDropdownProps): React.ReactElement => {
  const [isMenuShowed, setMenuShowed] = useState<boolean>(false);

  const [buttonValue, setButtonValue] = useState<string>(dropdown.initialValue);

  function handleMenuOnClick(val: string) {
    setMenuShowed(false);

    dropdown.setVal(val);

    setButtonValue(val);
  }

  return (
    <S.DropdownContainer>
      <button onClick={() => setMenuShowed(!isMenuShowed)}>
        <p>{buttonValue}</p>
        <S.Img
          $direction={isMenuShowed ? 'up' : 'down'}
          alt=""
          src={dropdownArrorIcon}
        ></S.Img>
      </button>

      {isMenuShowed && (
        <div>
          {dropdown.menus.map((menu) => (
            <button key={menu} onClick={() => handleMenuOnClick(menu)}>
              {menu}
            </button>
          ))}
        </div>
      )}
    </S.DropdownContainer>
  );
};

export default TopupDropdown;
