from typing import Annotated
from fastapi import Depends
from sqlmodel import SQLModel, Session, create_engine

from app.config import settings

engine = create_engine(url=settings.DB_URL.unicode_string())


def create_db_and_tables():
    SQLModel.metadata.create_all(engine)


def get_session():
    with Session(engine) as session:
        yield session


SessionDep = Annotated[Session, Depends(get_session)]
