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

const EditButton = ({ id, userKey, onEditSuccess, defaultTime }) => {
  const [open, setOpen] = useState(false);
  const [editedTime, setEditedTime] = useState("");

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
      <IconButton aria-label="edit" onClick={handleClickOpen} size="large">
        <EditIcon sx={{ fontSize: 30 }} />
      </IconButton>

      <Dialog
        open={open}
        onClose={handleClose}
        maxWidth="sm"
        fullWidth // ダイアログの幅を最大にします
        sx={{
          "& .MuiDialog-paper": {
            width: "100%", // ダイアログの幅を100%に設定
            maxWidth: "400px", // 最大幅を400pxに設定
          },
        }}
      >
        <DialogTitle>時間を編集</DialogTitle>
        <DialogContent
          sx={{
            display: "flex", // コンテンツをフレックスボックスにします
            flexDirection: "column", // アイテムを縦に並べます
          }}
        >
          <TextField
            autoFocus
            margin="dense"
            id="time"
            label="時間"
            type="time"
            fullWidth // 入力フィールドの幅を最大にします
            InputLabelProps={{
              shrink: true, // ラベルを縮小して時間が表示されるようにします
            }}
            variant="outlined" // アウトラインスタイルのTextFieldに変更
            value={editedTime}
            onChange={(e) => setEditedTime(e.target.value)}
            sx={{
              marginBottom: 2, // 下部のマージンを追加
            }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>キャンセル</Button>
          <Button onClick={handleEdit}>送信</Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default EditButton;
