import {check} from 'k6';
import {BeginObjectRPC} from 'k6/x/stbb';

var grant = "TODO"
var test = BeginObjectRPC(grant, "sj://qwe/k1")
test.init()

export default function () {
    var res = test.run()
    check(res, {
        "no error": (res) => {
            return res == null
        }
    })
}
