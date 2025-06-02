from datetime import datetime, timezone
import uuid
from pydantic import ConfigDict, field_validator
from sqlmodel import Column, DateTime, Field, SQLModel

from app.utils import validate_empty_str_to_none


class FormulaBase(SQLModel):
    model_config = ConfigDict(extra="forbid")  # type: ignore

    title: str = Field(min_length=1, max_length=255)
    description: str | None = None
    content: str = Field(min_length=1)

    @field_validator("description", mode="before")
    @classmethod
    def empty_str_to_none(cls, v: str | None):
        return validate_empty_str_to_none(v)


class Formula(FormulaBase, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    created_at: datetime = Field(
        default_factory=lambda: datetime.now(timezone.utc),
        sa_column=Column(DateTime(timezone=True)),
    )
    updated_at: datetime = Field(
        default_factory=lambda: datetime.now(timezone.utc),
        sa_column=Column(
            DateTime(timezone=True), onupdate=lambda: datetime.now(timezone.utc)
        ),
    )


class FormulaCreate(FormulaBase):
    pass


class FormulaUpdate(SQLModel):
    model_config = ConfigDict(extra="forbid")  # type: ignore

    title: str | None = Field(default=None, min_length=1, max_length=255)
    description: str | None = None
    content: str | None = Field(default=None, min_length=1)

    @field_validator("description", mode="before")
    @classmethod
    def empty_str_to_none(cls, v: str | None):
        return validate_empty_str_to_none(v)
