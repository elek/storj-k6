import {check} from 'k6';
import {UplinkTest} from 'k6/x/stbb';

var test = UplinkTest("__TODO__")

export default function () {
    var res = test()
    check(res, {
        "no error": (res) => {
            return res == null
        }
    })

}
