<?php

namespace App\Repositories;

use App\DTO\Date\DatePeriodDTO;

abstract class BaseRepository
{
    abstract protected function getModel();
    abstract protected function getResource();

    public function findAll(): array
    {
        $itens = $this->getModel()::orderBy('id', 'desc')->get();
        return $itens ? $this->getResource()->arrayToDtoItens($itens->toArray()) : array();
    }

    public function findById(int $id)
    {
        $item = $this->getModel()->find($id);
        return $item ? $this->getResource()->arrayToDto($item->toArray()) : null;
    }

    public function insert($item)
    {
        $array = $this->getResource()->dtoToArray($item);
        $inserted = $this->getModel()->create($array)->toArray();
        return $this->getResource()->arrayToDto($inserted);
    }

    public function update(int $id, $item)
    {
        $array = $this->getResource()->dtoToArray($item);
        $this->getModel()->where('id', $id)->update($array);
        return $this->findById($id);
    }

    public function deleteById(int $id): bool
    {
        $item = $this->getModel()->find($id);
        $item?->delete();
        return true;
    }

    public function findByPeriod(DatePeriodDTO $period): array
    {
        $itens = $this->getModel()
            ->select()
            ->where('created_at', '>=', $period->getStartDate())
            ->where('created_at', '<=', $period->getEndDate())
            ->orderBy('id', 'desc')
            ->get();
        return $this->getResource()->arrayToDtoItens($itens->toArray());
    }
}