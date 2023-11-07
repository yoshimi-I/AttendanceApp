import React, { useState } from 'react';
import { useRouter } from "next/navigation";
import { Dialog, DialogActions, DialogTitle, Button, IconButton } from '@mui/material';
import LogoutIcon from '@mui/icons-material/Logout';
import { getAuth, signOut } from 'firebase/auth';

const LogoutDialog = () => {
  const [open, setOpen] = useState(false);
  const router = useRouter();
  const auth = getAuth();

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleLogout = async () => {
    try {
      await signOut(auth);
      handleClose();
      router.push('/signin');
    } catch (error) {
      console.error('Logout failed:', error);
    }
  };

  return (
    <>
      <IconButton onClick={handleClickOpen}>
        <LogoutIcon />
      </IconButton>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>ログアウトしますか？</DialogTitle>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            キャンセル
          </Button>
          <Button onClick={handleLogout} color="primary" autoFocus>
            ログアウト
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default LogoutDialog;
