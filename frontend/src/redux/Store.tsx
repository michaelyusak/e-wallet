import { configureStore } from '@reduxjs/toolkit';
import userSlice from '../features/UserData/userSlice';

const store = configureStore({
  reducer: {
    user: userSlice.reducer,
  },
});

export default store;
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
