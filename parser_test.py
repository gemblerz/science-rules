import unittest

from parser import evaluate

show_parse_tree = False

class TestEvaluate(unittest.TestCase):
    def test_simle_rules(self):
        tests = {
            "test1": {
                "rule": "1 + 2 == 3",
                "want": True,
            },
            "test2": {
                "rule": "(1 + 2) * 2 == 6",
                "want": True,
            },
            "test3": {
                "rule": "(1 + 2) * 2 == 6 and 1 + 2 == 4",
                "want": False,
            },
            "test4": {
                "rule": "false == true",
                "want": False,
            },
            "test5": {
                "rule": "env.event.smoke == true",
                "want": True,
            },
            # "test6": {
            #     "rule": "sum(env.temperature{sensor=bme680}[-5m]) > 32",
            #     "want": True,
            # },
        }
        for k in tests:
            t = tests.get(k)
            rule = t.get("rule")
            want = t.get("want")
            result = evaluate(rule, {}, show_tree=show_parse_tree)
            self.assertEqual(result, want, msg=f'{k}: {rule} expected {want}, but got {result}')

if __name__ == '__main__':
    unittest.main()