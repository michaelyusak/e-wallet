import { IResponse } from './APIResponse';
import { IWallet } from './Wallet';

interface User {
  id: string;
  email: string;
  name: string;
  profile_picture: string;
}

export interface IUserDetail extends User {
  wallet: IWallet;
}

export interface IUserDetailResponse extends IResponse {
  data: IUserDetail;
}
