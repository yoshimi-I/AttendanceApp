import { ActivityCalendar } from "@/components/githubCarender/githubCarender";
import CurrentTime from "@/components/time/currentTime";


const MyPage = () => {
  const sampleData = [
    {
      day: "2023-01-01",
      activity: 5
    },
    {
      day: "2023-11-02",
      activity: 1
    }
  ]
  return (
    <main>
      <CurrentTime />
      <ActivityCalendar sampleData={sampleData} />
    </main>
  );
}
export default MyPage;
