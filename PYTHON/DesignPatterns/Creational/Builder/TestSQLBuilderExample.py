from SQLBuilderExample import *
import unittest

class TestPostgresqlQueryBuilder(unittest.TestCase):
    def setUp(self) -> None:
        self._table_name = "tests"
        self._limit_keyword = "LIMIT"
        self.sql_builder = PostgresqlQueryBuilder(True)
    
    def test_select_all(self):
        query_res = "SELECT * FROM %s" % self._table_name
        
        self.assertEqual(
            self.sql_builder.select(
            self._table_name
        ).get_query(),
            query_res
        )
    
    def test_select_not_all(self):
        fields_name = ["names", "families", "ages"]
        query_res = "SELECT %s FROM %s" % (
            "names, families, ages", self._table_name
        )
        self.assertEqual(
            self.sql_builder.select(
                self._table_name, fields_name
            ).get_query(),
            query_res
        )

class TestOracleQueryBuilder(TestPostgresqlQueryBuilder):
    def setUp(self) -> None:
        self._table_name = "tests"
        self.sql_builder = OracleQueryBuilder()

if __name__ == '__main__':
    unittest.main()