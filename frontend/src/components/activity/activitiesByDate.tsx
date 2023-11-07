import React from "react";
import {
  Container,
  Paper,
  Typography,
  Box,
  List,
  ListItem,
  ListItemText,
  Divider,
} from "@mui/material";

interface Activity {
  id: number;
  type: string;
  start_time: string;
  end_time: string;
}

interface ActivitiesComponentProps {
  data: {
    date: string;
    activities: Activity[];
  };
}

const ActivitiesByDate: React.FC<ActivitiesComponentProps> = ({ data }) => {
  const { date, activities } = data;

  return (
    <Container maxWidth="sm" sx={{ mt: 4, mb: 4 }}>
      <Box
        sx={{
          my: 4,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography
          variant="h3"
          component="h1"
          gutterBottom
          sx={{
            fontWeight: "bold", // フォントの太さ
            color: "secondary.main", // テーマのセカンダリーカラーを使用
            fontFamily: "Courier New", // カーシブ体のフォントファミリー
            textAlign: "center", // 中央揃え
            mt: 2, // 上のマージン
            mb: 4, // 下のマージン
            textShadow: "0px 2px 3px rgba(0,0,0,0.3)", // テキストシャドウを追加
          }}
        >
          {date}
        </Typography>

        {activities && activities.length > 0 ? (
          <List sx={{ width: "100%" }}>
            {activities.map((activity, index) => (
              <React.Fragment key={activity.id}>
                <Paper elevation={2} sx={{ mb: 2, overflow: "hidden" }}>
                  <ListItem sx={{ justifyContent: "center" }}>
                    <ListItemText
                      primary={
                        <Box
                          sx={{
                            textAlign: "center",
                            width: "fit-content",
                            margin: "0 auto",
                          }}
                        >
                          <Typography
                            variant="h6"
                            sx={{
                              fontWeight: "bold",
                              color:
                                activity.type === "作業"
                                  ? "primary.main"
                                  : activity.type === "休憩"
                                  ? "secondary.main"
                                  : "error.main",
                              backgroundColor:
                                activity.type === "作業"
                                  ? "#e3f2fd"
                                  : activity.type === "休憩"
                                  ? "#fce4ec"
                                  : "#fffde7", // 背景色を設定
                              borderRadius: "10px",
                              padding: "8px 60px",
                              display: "inline-block",
                              boxShadow: "0 3px 5px rgba(0, 0, 0, 0.2)",
                              mt: 2,
                              mb: 2,
                            }}
                          >
                            {activity.type}
                          </Typography>
                        </Box>
                      }
                      secondary={
                        <>
                          <Box sx={{ textAlign: "center", mt: 3, mb: 2 }}>
                            <Typography
                              component="span"
                              variant="h6" // より大きなサイズに変更
                              color="text.secondary"
                              sx={{ display: "block" }} // fontWeightの調整は削除
                            >
                              開始時間: {activity.start_time}
                            </Typography>
                          </Box>
                          <Box sx={{ textAlign: "center", mt: 2, mb: 3 }}>
                            <Typography
                              component="span"
                              variant="h6" // より大きなサイズに変更
                              color="text.secondary"
                              sx={{ display: "block" }} // fontWeightの調整は削除
                            >
                              終了時間: {activity.end_time}
                            </Typography>
                          </Box>
                        </>
                      }
                    />
                  </ListItem>
                </Paper>
                {index < activities.length - 1 && <Divider />}
              </React.Fragment>
            ))}
          </List>
        ) : (
          <Typography variant="body1">データが存在しません。</Typography>
        )}
      </Box>
    </Container>
  );
};

export default ActivitiesByDate;
