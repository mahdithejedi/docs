from abc import ABC, abstractmethod


class QueryBuilder(ABC):

    def __init__(self, allow_multiple_query: bool=False) -> None:
        self._allow_multiple_query = allow_multiple_query
        self._reset()

    def reset(self):
        if self._allow_multiple_query and self._query:
            self._query = '%s;' % self._query
        else:
            self._reset()

    def _reset(self):
        self._query = str()

    @abstractmethod
    def select(self, table: str, fields: list = []) -> str:
        pass

    @abstractmethod
    def where(self, where: dict = {}) -> str:
        pass

    @abstractmethod
    def update(self, table: str, sets: dict = {}) -> str:
        pass

    @abstractmethod
    def limit(self, limit_count: int = 0) -> str:
        pass

    def get_query(self, force_reset: bool = True) -> str:
        query = self._query
        if force_reset:
            self._reset()
        return query


class PostgresqlQueryBuilder(QueryBuilder):
    def select(self, table: str, fields: list = []) -> str:
        self.reset()
        self._query += "select %s from %s" % (
            ', '.join(fields) if fields else '*'
            , table
        )
        return self
    
    def where(self, where: dict = {}) -> str:
        self._query += " where "
        for key, value in where.items():
            self._query += ' %s=%s' %(key, value)
            self._query += ","
        self._query = self._query[:-1]
        return self

    def update(self, table: str, sets: dict = {}) -> str:
        self.reset()
        self._query += " update %s set" % table
        for key, value in sets.items():
            self._query += " %s=%s" %(key, value)
            self._query += ","
        self._query = self._query[:-1]
        return self

    def limit(self, limit_count: int = 0) -> str:
        self._query += " LIMIT %s" % limit_count
        return self


class OracleQueryBuilder(PostgresqlQueryBuilder):
    def limit(self, limit_count: int = 0) -> str:
        self._query += " OFFSET %s" % limit_count()
        return self


class ORM:
    def __init__(self, query_builder: QueryBuilder) -> None:
        self._query_builder = query_builder(True)

    def real_select(self, table: str, fields: list = [], where: dict = {}):
        return self._query_builder.select(table, fields).where(where).get_query()

    def real_update(self, table: str, sets: dict = {}, where: dict = {}):
        return self._query_builder.update(table, sets).where(where).get_query()

if __name__ == '__main__':
    orm = ORM(OracleQueryBuilder)
    from time import time
    print(
        orm.real_update(
            "user_attrs", {
                "finish_credit_threshold": int(time()),
                "real_time_login": 0
            },
            {
                "user_id": 89,
                "real_time_login": int(time())
            }
        )
    )