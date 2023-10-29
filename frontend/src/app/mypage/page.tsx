import { ActivityCalendar } from "@/components/githubCarender/githubCarender";
import TaskButtons from "@/components/register/registerButton";
import CurrentTime from "@/components/time/currentTime";


const MyPage = () => {
  const sampleData = [
    {
      day: "2023-01-01",
      activity_time: 5
    },
    {
      day: "2023-11-02",
      activity_time: 1
    }
  ]
  return (
    <main>
      <CurrentTime />
      <TaskButtons />
      <ActivityCalendar sampleData={sampleData} />
    </main>
  );
}
export default MyPage;
