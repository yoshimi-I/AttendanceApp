import React, { useState } from "react";
import axios from "axios";
import IconButton from "@mui/material/IconButton";
import EditIcon from "@mui/icons-material/Edit";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { TimeToString } from "../../../util/timeToString";

const EditButton = ({ id, userKey, onEditSuccess, defaultTime }) => {
  const [open, setOpen] = useState(false);
  const [editedTime, setEditedTime] = useState(TimeToString(defaultTime));

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleEdit = async () => {
    try {
      const response = await axios.post("/api/edit-time", {
        id: id,
        user_key: userKey,
        time: editedTime,
      });

      onEditSuccess(response.data);
      handleClose();
    } catch (error) {
      console.error("Error updating time:", error);
    }
  };

  return (
    <>
      <IconButton aria-label="edit" onClick={handleClickOpen} size="large" >
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
            maxWidth: "400px",
            padding: "20px",
            backgroundColor: "#F5F5F5",
          },
        }}
      >
        <DialogTitle sx={{ color: "#333", fontWeight: "bold" }}>時間を編集</DialogTitle>
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
              '& .MuiOutlinedInput-root': {
                borderRadius: '8px',
              },
            }}
          />
        </DialogContent>
        <DialogActions sx={{ justifyContent: "space-between" }}>
          <Button onClick={handleClose} variant="outlined" color="secondary">キャンセル</Button>
          <Button onClick={handleEdit} variant="contained" color="primary">送信</Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default EditButton;
