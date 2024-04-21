import { PayloadAction, createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { IUserDetail } from '../../interfaces/User';
import Cookies from 'js-cookie';

type UserState = {
  user: IUserDetail;
  isLoading: boolean;
  error: string;
};

const initialState: UserState = {
  user: {
    name: '',
    email: '',
    id: '',
    profile_picture: '',
    wallet: {
      balance: '',
      id: '',
      income: '',
      expense: '',
      number: '',
    },
  },
  isLoading: false,
  error: '',
};

export const fetchUser = createAsyncThunk('user/fetch', async () => {
  const url = 'http://localhost:8080/users';
  const token = Cookies.get('token') ?? '';
  const options = {
    method: 'GET',
    headers: { 'Content-Type': 'application/json', Authorization: token },
  };

  const response = await fetch(url, options);
  const responseData = await response.json();

  console.log('trigerred')

  return responseData.data;

});

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(
      fetchUser.fulfilled,
      (state, action: PayloadAction<IUserDetail>) => {
        state.isLoading = false;
        state.user = action.payload;
        state.error = '';
      },
    );
    builder.addCase(fetchUser.pending, (state) => {
      state.isLoading = true;
      state.error = '';
    });
    builder.addCase(fetchUser.rejected, (state, action) => {
      state.isLoading = false;
      state.error = action.error.message || 'failed to fetch user detail';
    });
  },
});

export default userSlice;
