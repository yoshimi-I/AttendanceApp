import { ActivityCalendar } from "@/components/githubCarender/githubCarender";
import CurrentTime from "../components/time/currentTime";
import TaskButtons from "@/components/register/registerButton";
import AttendanceButton from "@/components/register/registerButton";
import AttendanceButtons from "@/components/register/registerButtonList";
export default function Home() {
  const sampleData = [
    {
      day: "2023-01-01",
      activity_time: 5,
    },
    {
      day: "2023-11-02",
      activity_time: 1,
    },
  ];
  type ButtonType = "作業開始" | "作業終了" | "休憩開始" | "休憩終了";
  const handleActivityAction = (type: ButtonType) => {
    // ここで何らかの処理を行います。例えば、API呼び出しや状態の更新など
    console.log(`${type} ボタンがクリックされました`);
  };
  return (
    <main>
      <CurrentTime />
      <AttendanceButtons />
      <ActivityCalendar sampleData={sampleData} />
    </main>
  );
}
