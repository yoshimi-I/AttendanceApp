"use client";
import { Box } from "@mui/system";
import AttendanceButton from "./registerButton";
import { Paper, Typography } from "@mui/material";
import { useState } from "react";

type AttendanceButtonsProps = {
  userStatus: string;
};

const AttendanceButtons: React.FC<AttendanceButtonsProps> = ({
  userStatus: initialUserStatus,
}) => {
  const [status, setStatus] = useState(initialUserStatus);

  return (
    <>
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        style={{ marginBottom: "20px" }}
      >
        <Paper
          elevation={3}
          style={{
            width: "400px",
            height: "120px",
            borderRadius: "100px",
            color: "#862fa8",
            textAlign: "center",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Typography
            variant="h3"
            style={{ fontWeight: "bold", fontFamily: "Arial" }}
          >
            {status}
          </Typography>
        </Paper>
      </Box>
      <Box
        display="flex"
        flexDirection="row"
        gap={2}
        alignItems="center"
        justifyContent="center"
        style={{ marginBottom: "40px" }}
      >
        <AttendanceButton type="作業開始" userStatus={setStatus} />
        <AttendanceButton type="作業終了" userStatus={setStatus} />
        <AttendanceButton type="休憩開始" userStatus={setStatus} />
        <AttendanceButton type="休憩終了" userStatus={setStatus} />
      </Box>
    </>
  );
};

export default AttendanceButtons;
