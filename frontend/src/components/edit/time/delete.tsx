import { useState } from "react";
import IconButton from "@mui/material/IconButton";
import DeleteIcon from "@mui/icons-material/Delete";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { Alert, AlertColor, Button, Snackbar } from "@mui/material";

const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

type DeleteButtonProps = {
  id: number;
  userKey: string;
  onDeleteSuccess: () => void;
};

const DeleteButton: React.FC<DeleteButtonProps> = ({
  id,
  userKey,
  onDeleteSuccess,
}) => {
  const [open, setOpen] = useState(false);
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState("");
  const [snackbarSeverity, setSnackbarSeverity] = useState<AlertColor>("error");

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleDelete = async () => {
    try {
      const response = await fetch(`${baseUrl}/study/activity/delete`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          activity_id: id,
          user_key: userKey,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        setSnackbarMessage(errorData.error || "不明なエラーが発生しました");
        setSnackbarSeverity("error");
      } else {
        setSnackbarMessage("削除に成功しました");
        setSnackbarSeverity("success");
        handleClose();
        onDeleteSuccess();
      }
    } catch (error) {
      console.error("Error deleting time:", error);
      setSnackbarMessage("エラーが発生しました");
      setSnackbarSeverity("error");
    } finally {
      setOpenSnackbar(true);
    }
  };

  const handleCloseSnackbar = () => {
    setOpenSnackbar(false);
  };

  return (
    <>
      <IconButton aria-label="delete" onClick={handleClickOpen} size="large">
        <DeleteIcon sx={{ fontSize: 30 }} />
      </IconButton>

      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
        sx={{
          "& .MuiDialog-paper": {
            minWidth: "400px",
            backgroundColor: "#f5f5f5",
          },
        }}
      >
        <DialogTitle id="alert-dialog-title">時間を削除</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            本当に削除しますか？
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary" variant="outlined">
            キャンセル
          </Button>
          <Button
            onClick={handleDelete}
            color="primary"
            variant="contained"
            autoFocus
          >
            削除
          </Button>
        </DialogActions>
      </Dialog>
      <Snackbar
        open={openSnackbar}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert
          onClose={handleCloseSnackbar}
          severity={snackbarSeverity}
          sx={{ width: "100%" }}
        >
          {snackbarMessage}
        </Alert>
      </Snackbar>
    </>
  );
};

export default DeleteButton;
