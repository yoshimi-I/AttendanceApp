import { useState } from "react";
import IconButton from "@mui/material/IconButton";
import EditIcon from "@mui/icons-material/Edit";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { TimeToString } from "../../../util/timeToString";
import { Alert, AlertColor, Snackbar } from "@mui/material";
import { StrToJSDate } from "../../../util/japanTime";

const EditButton = ({ id, userKey, onEditSuccess, defaultTime }) => {
  const [open, setOpen] = useState(false);
  const [editedTime, setEditedTime] = useState(TimeToString(defaultTime));
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState("");
  const [snackbarSeverity, setSnackbarSeverity] = useState<AlertColor>("error");

  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleEdit = async () => {
    try {
      const response = await fetch(`${baseUrl}/study/activity/update`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          activity_id: id,
          user_key: userKey,
          time: StrToJSDate(editedTime),
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        setSnackbarMessage(errorData.error || "不明なエラーが発生しました");
        setSnackbarSeverity("error");
      } else {
        setSnackbarMessage("修正に成功しました");
        setSnackbarSeverity("success");
        handleClose();
        onEditSuccess();
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
      <IconButton aria-label="edit" onClick={handleClickOpen} size="large">
        <EditIcon sx={{ fontSize: 30 }} />
      </IconButton>

      <Dialog
        open={open}
        onClose={handleClose}
        maxWidth="sm"
        fullWidth
        sx={{
          "& .MuiDialog-paper": {
            width: "100%",
            maxWidth: "300px",
            padding: "20px",
            backgroundColor: "#F5F5F5",
          },
        }}
      >
        <DialogTitle sx={{ color: "#333", fontWeight: "bold" }}>
          時間を編集
        </DialogTitle>
        <DialogContent
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: "20px",
          }}
        >
          <TextField
            autoFocus
            margin="dense"
            id="time"
            label="時間"
            type="time"
            fullWidth
            InputLabelProps={{
              shrink: true,
            }}
            variant="outlined"
            value={editedTime}
            onChange={(e) => setEditedTime(e.target.value)}
            sx={{
              "& .MuiOutlinedInput-root": {
                borderRadius: "8px",
              },
            }}
          />
        </DialogContent>
        <DialogActions sx={{ justifyContent: "space-between" }}>
          <Button onClick={handleClose} variant="outlined" color="primary">
            キャンセル
          </Button>
          <Button onClick={handleEdit} variant="contained" color="primary">
            送信
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

export default EditButton;
