import uuid
from fastapi import APIRouter, HTTPException
from sqlmodel import select

from app.database import SessionDep
from app.models.formula import Formula, FormulaCreate, FormulaUpdate

router = APIRouter(tags=["formulas"])


@router.post("/formulas")
def create_formula(formula: FormulaCreate, session: SessionDep) -> Formula:
    db_formula = Formula.model_validate(formula)
    session.add(db_formula)
    session.commit()
    session.refresh(db_formula)
    return db_formula


@router.get("/formulas/{id}")
def read_formula(id: uuid.UUID, session: SessionDep) -> Formula:
    formula = session.get(Formula, id)
    if not formula:
        raise HTTPException(status_code=404, detail="Formula not found")
    return formula


@router.patch("/formulas/{id}")
def update_formula(
    id: uuid.UUID, formula: FormulaUpdate, session: SessionDep
) -> Formula:
    db_formula = session.get(Formula, id)
    if not db_formula:
        raise HTTPException(status_code=404, detail="Formula not found")
    formula_data = formula.model_dump(exclude_unset=True)
    db_formula.sqlmodel_update(formula_data)
    session.add(db_formula)
    session.commit()
    session.refresh(db_formula)
    return db_formula


@router.delete("/formulas/{id}")
def delete_formula(id: uuid.UUID, session: SessionDep):
    formula = session.get(Formula, id)
    if not formula:
        raise HTTPException(status_code=404, detail="Formula not found")
    session.delete(formula)
    session.commit()
    return {"ok": True}


@router.get("/formulas", response_model=list[Formula])
def read_formulas(session: SessionDep):
    formulas = session.exec(select(Formula)).all()
    return formulas
