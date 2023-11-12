import {
  Container,
  Paper,
  Typography,
  List,
  ListItem,
  ListItemText,
  Box,
} from "@mui/material";
import { Fragment } from "react";
import SchoolIcon from "@mui/icons-material/School";
import CoffeeIcon from "@mui/icons-material/LocalCafe";
import EditButton from "../edit/time/edit";
import { TimeToString } from "../../util/timeToString";
import DeleteButton from "../edit/time/delete";
import { User } from "firebase/auth";

interface Activity {
  id: number;
  type: string;
  time: Date;
}

interface ActivitiesComponentProps {
  data: {
    date: string;
    activities: Activity[];
  };
  user: User | null | undefined;
  refreshData: (user: User) => void;
}

const ActivitiesByDate: React.FC<ActivitiesComponentProps> = ({
  data,
  user,
  refreshData,
}) => {
  const { date, activities } = data;

  if (!user) {
    console.error("User is null or undefined.");
    return;
  }

  const handleSuccess = () => {
    refreshData(user); // データを更新
  };

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
          {activities.map((activity) => (
            <Fragment key={activity.id}>
              <Paper
                elevation={2}
                sx={{
                  mb: 2,
                  width: "100%",
                  overflow: "hidden",
                  borderRadius: "20px",
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
                          width: "100%",
                        }}
                      >
                        <Box
                          sx={{
                            flex: 4,
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
                            {TimeToString(activity.time)}
                          </Typography>
                        </Box>
                        <Box
                          sx={{
                            flex: 1, // 全体の1割のスペースを取る
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "center", // 中央寄せ
                          }}
                        >
                          <EditButton
                            id={activity.id}
                            userKey={user.uid}
                            defaultTime={activity.time}
                            onEditSuccess={handleSuccess}
                          />
                        </Box>
                        <Box
                          sx={{
                            flex: 1,
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "center",
                          }}
                        >
                          <DeleteButton
                            id={activity.id}
                            userKey={user.uid}
                            onDeleteSuccess={handleSuccess}
                          />
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
