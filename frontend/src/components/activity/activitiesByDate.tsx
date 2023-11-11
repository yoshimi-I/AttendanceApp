import {
  Container,
  Paper,
  Typography,
  List,
  ListItem,
  ListItemText,
  Divider,
  Box,
  IconButton,
} from "@mui/material";
import { Fragment } from "react";
import SchoolIcon from "@mui/icons-material/School";
import CoffeeIcon from "@mui/icons-material/LocalCafe";
import EditIcon from "@mui/icons-material/Edit";

interface Activity {
  id: number;
  type: string;
  time: string;
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
      <Typography
        variant="h3"
        component="h1"
        gutterBottom
        sx={{
          fontWeight: "bold",
          color: "secondary.main",
          fontFamily: "Courier New",
          textAlign: "center",
          mt: 2,
          mb: 4,
          textShadow: "0px 2px 3px rgba(0,0,0,0.3)",
        }}
      >
        {date}
      </Typography>

      {activities && activities.length > 0 ? (
        <List sx={{ width: "100%" }}>
          {activities.map((activity, index) => (
            <Fragment key={activity.id}>
              <Paper
                elevation={2}
                sx={{
                  mb: 2,
                  overflow: "hidden",
                  borderRadius: "20px", // PaperのborderRadiusも調整
                  boxShadow: "4",
                  padding: "5px",
                  margin: "20px",
                }}
              >
                <ListItem sx={{ justifyContent: "flex-start" }}>
                  <ListItemText
                    primary={
                      <Box
                        sx={{
                          display: "flex",
                          alignItems: "center",
                          width: "100%", // 親要素の幅に合わせる
                        }}
                      >
                        <Box
                          sx={{
                            flex: 4, // 全体の4割のスペースを取る
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "flex-start",
                            backgroundColor:
                              activity.type === "作業開始" ||
                              activity.type === "作業終了"
                                ? "#e3f2fd"
                                : "#fce4ec",
                            borderRadius: "20px",
                            padding: "8px 20px",
                            boxShadow: "0 3px 5px rgba(0, 0, 0, 0.2)",
                          }}
                        >
                          {activity.type.includes("作業") && (
                            <SchoolIcon
                              sx={{
                                fontSize: 30,
                                mr: 2,
                                color:
                                  activity.type === "作業開始" ||
                                  activity.type === "作業終了"
                                    ? "primary.main"
                                    : "secondary.main",
                              }}
                            />
                          )}
                          {activity.type.includes("休憩") && (
                            <CoffeeIcon
                              sx={{
                                fontSize: 30,
                                mr: 2,
                                color:
                                  activity.type === "作業開始" ||
                                  activity.type === "作業終了"
                                    ? "primary.main"
                                    : "secondary.main",
                              }}
                            />
                          )}
                          <Typography
                            variant="h6"
                            sx={{
                              fontWeight: "bold",
                              textAlign: "center",
                              color:
                                activity.type === "作業開始" ||
                                activity.type === "作業終了"
                                  ? "primary.main"
                                  : "secondary.main",
                            }}
                          >
                            {activity.type}
                          </Typography>
                        </Box>

                        <Box
                          sx={{
                            flex: 7, // 全体の6割のスペースを取る
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "center", // 右寄せ
                          }}
                        >
                          <Typography
                            component="span"
                            variant="h4"
                            sx={{
                              fontWeight: "bold",
                              color: "text.secondary",
                            }}
                          >
                            {activity.time}
                          </Typography>
                        </Box>
                        <Box
                          sx={{
                            flex: 2, // 全体の1割のスペースを取る
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "center", // 中央寄せ
                          }}
                        >
                          <IconButton
                            aria-label="edit"
                            onClick={() => {
                              console.log("編集ボタンがクリックされました。");
                            }}
                          >
                            <EditIcon sx={{ fontSize:30 }} />
                          </IconButton>
                        </Box>
                      </Box>
                    }
                  />
                </ListItem>
              </Paper>
            </Fragment>
          ))}
        </List>
      ) : (
        <Typography variant="body1">データが存在しません。</Typography>
      )}
    </Container>
  );
};

export default ActivitiesByDate;
