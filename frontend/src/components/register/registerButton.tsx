import { useState } from "react";
import Button from "@mui/material/Button";
import { getAuth } from "firebase/auth";
import Snackbar from "@mui/material/Snackbar";
import Alert, { AlertColor } from "@mui/material/Alert";
import { getJSTDate } from "../../hooks/japanTime";

type ButtonType = "作業開始" | "作業終了" | "休憩開始" | "休憩終了";
const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

interface AttendanceButtonProps {
  type: ButtonType;
  userStatus: (newStatus: string) => void;
}

const AttendanceButton: React.FC<AttendanceButtonProps> = ({
  type,
  userStatus,
}) => {
  let color: "primary" | "secondary" | "inherit";
  let endpoint: string;
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState("");
  const [snackbarSeverity, setSnackbarSeverity] = useState<AlertColor>("error");

  switch (type) {
    case "作業開始":
      color = "primary";
      endpoint = "/study/activity/work/start";
      break;
    case "作業終了":
      color = "secondary";
      endpoint = "/study/activity/work/end";
      break;
    case "休憩開始":
      color = "inherit";
      endpoint = "/study/activity/break/start";
      break;
    case "休憩終了":
      color = "inherit";
      endpoint = "/study/activity/break/end";
      break;
    default:
      return null;
  }

  const handleClick = async () => {
    const auth = getAuth();
    const user = auth.currentUser;
    const userKey = user ? user.uid : null;

    if (!userKey) {
      setSnackbarMessage("ユーザーがログインしていません。");
      setSnackbarSeverity("error");
      setOpenSnackbar(true);
      return;
    }

    const currentTime = getJSTDate();
    console.log(currentTime);
    const url = `${baseUrl}${endpoint}`;

    try {
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_key: userKey,
          time: currentTime,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        setSnackbarMessage(errorData.error || "不明なエラーが発生しました");
        setSnackbarSeverity("error");
      } else {
        const responseData = await response.json();
        setSnackbarMessage("登録に成功しました");
        setSnackbarSeverity("success");
        userStatus(responseData.status);
      }
    } catch (error) {
      setSnackbarMessage("エラーが発生しました: " + error.message);
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
      <Button
        variant="contained"
        color={color}
        size="large"
        sx={{
          borderRadius: "50%",
          width: "150px",
          height: "150px",
          margin: "2%",
        }}
        onClick={handleClick}
      >
        {type}
      </Button>
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

export default AttendanceButton;
