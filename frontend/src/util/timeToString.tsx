import { ToJSDate } from "./japanTime";

export const  TimeToString = (dateInput: Date | string | number) =>  {
  const date = ToJSDate(dateInput);
  const hours = date.getHours().toString().padStart(2, '0');
  const minutes = date.getMinutes().toString().padStart(2, '0');
  return `${hours}:${minutes}`;
}

export const TimeSplit = (dateInput: Date | string | number) => {
  const date = ToJSDate(dateInput);
  const hours = date.getHours();
  const minutes = date.getMinutes();
  return [hours, minutes];
}
