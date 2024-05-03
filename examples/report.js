import stbb from 'k6/x/stbb';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

var s = stbb("cockroach://root@localhost:26257/metainfo?sslmode=disable")
  
export default function () {
  s()
}

export function handleSummary(data) {
  return {
    "summary.html": htmlReport(data),
  };
}
