import { ToJSDate } from "./japanTime";

function TimeToString(dateInput: Date | string | number): string {
  const date = ToJSDate(dateInput);
  const hours = date.getHours().toString().padStart(2, '0');
  const minutes = date.getMinutes().toString().padStart(2, '0');
  return `${hours} : ${minutes}`;
}

export default TimeToString;
